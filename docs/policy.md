# Policy

AlertChain has two types of policies: "Alert Policy" and "Action Policy". Both policies are written in the Rego language. This document describes the input and output schemas for these policies. See [Open Policy Agent official document](https://www.openpolicyagent.org/docs/latest/policy-language/) for more detail of Rego language.

The **Alert Policy** is responsible for determining whether the incoming event data from external sources should be treated as an alert or not. For example, when receiving notifications from external services, you may want to handle only alerts related to specific categories, or you may want to exclude events that meet certain conditions (such as specific users or hosts). The Alert Policy can be used to achieve these goals by excluding certain events or including only specific events as alerts.

On the other hand, the **Action Policy** determines the appropriate response for detected alerts. For example, when an issue is detected on a cloud instance, the response may differ depending on the type of alert or the elements involved in the alert, such as stopping the instance, restricting the instance's communication, or notifying an administrator. You may also want to retrieve reputation information from external services and adjust the response accordingly. The Action Policy is responsible for defining and controlling these response procedures.

## Alert Policy

### Package

The package name for Alert Policy must follow the naming convention below:

```rego
package alert.{schema}
```

Here, `{schema}` must match the `{schema}` specified when receiving event data. For example, if the endpoint path for receiving data via Pub/Sub is `/alert/pubsub/my_alert`, the policy `package alert.my_alert` will be called.

### Input

The input for Alert Policy will be the structured data (mainly JSON) received from the previous phase. For example, if the following message is input via Google Cloud Pub/Sub:

```
{
    "message": {
        "data": "ewogICAgIm5hbWUiOiAic3VzcGljaW91c19hY3Rpb24iLAogICAgInVzZXIiOiAibS1taXp1dGFuaSIKfQo=",
    },
}
```

From the Pub/Sub schema, `message.data` is extracted, and `ewogICAgIm5hbWUiOiAic3VzcGljaW91c19hY3Rpb24iLAogICAgInVzZXIiOiAibS1taXp1dGFuaSIKfQo=` is Base64 decoded to:

```json
{
    "name": "suspicious_action",
    "user": "m-mizutani"
}
```

This data is stored in Rego's `input`. The policy will determine whether this data will be treated as an alert or not based on this data.

### Output

Once the alert determination is made, store the data with the schema below in the `alert` rule. The stored data will be treated as an alert. Output schema is according to Alert structure.

The `attrs` field (Attribute) serves not only to extract event data fields but also to accommodate user-defined values. For instance, users can add their own `severity` key Attribute to determine the appropriate action. Attributes bind the alert and can be added or replaced by the action policy. (Refer to the Action Policy section for more details)

### Example

```rego
package alert.my_alert

alert[msg] {
    input.name == "suspicious_action"
    msg := {
        "title": "detected suspicious action",
        "attrs": [
            {
                "key": "subject",
                "value": input.user,
            },
        ],
    }
}
```

In this example, the policy checks if the input contains the name "suspicious_action". If it does, an alert will be created with the title "detected suspicious_action" and a Attribute of "subject" set to "m-mizutani".

## Action Policy

An Action Policy is responsible for defining the following:

- `init` rule
  - Check alert data at first in action workflow
  - Abort all action processes if necessary
- `run` rule
  - Choose the next action to execute
  - Provide arguments for the next action
- `exit` rule
  - Create new or replace Attributes for the next action
  - Abort all action processes if necessary

The relationship between the `init`, `run` and `exit` rules in the Action Policy and the execution order of actions is illustrated in the diagram below.

![Action Policy Evaluation Workflow](https://github.com/m-mizutani/alertchain/assets/605953/b1c731e3-ceaa-4c77-a008-d4c0766c81a7)

When an alert is detected by the Alert Policy, the `init`, `run` and `exit` rules within the Action Policy are called alternately. The `init` rules can check alert data before calling `run` rules. The `run` rule can specify the execution of multiple actions. Also, the `exit` rule is called after each action is completed. If no actions are selected by the `run` rule, all operations will terminate.

```rego
package action

run[res] {
    input.seq == 0
    res := {
        "uses": "chatgpt.query",
        "args": {
            "secret_api_key": input.env.CHATGPT_API_KEY,
        },
    }
}
```

In this example, an action called `chatgpt.query` is launched to query the alert content to ChatGPT. The action to be launched is specified by `uses`, and the required arguments are specified by `args`. The `input.seq` value increments by 1 each time the `run` rule is called. Therefore, when the `run` rule is called for the second time, the result of `input.seq == 0` will be false, making the rule invalid, and no subsequent actions will be specified. If no actions are specified, the entire process will stop.

The `exit` rule primarily handles the transfer of Attributes obtained from the results of actions. Let's modify the ChatGPT calling rule slightly.

```rego
package action

run[res] {
    input.seq == 0
    res := {
        "id": "ask-chatgpt",
        "uses": "chatgpt.query",
        "args": {
            "secret_api_key": input.env.CHATGPT_API_KEY,
        },
    }
}

exit[res] {
    input.action.id == "ask-chatgpt"
    res := {
        "attrs": [
            {
                "key": "ChatGPT's comment",
                "value": input.action.result.choices[0].message.content,
            }
        ]
    }
}
```

We added the `id` value `ask-chatgpt` to the `run` rule, and then checked for it in the `exit` rule with `input.action.id == "ask-chatgpt"`. This ensures that the `exit` rule is only valid after the first `run` rule has been executed. In this `exit` rule, we extract the response message stored in the result of the ChatGPT action (https://platform.openai.com/docs/guides/chat/response-format) and store it as a Attribute. The stored Attribute will then be available for use in subsequent `run` and `exit` rules.

After the `exit` rule is called, the `run` rule is called again. For example, by adding the following `run` rule, we can ensure that the `run` rule is only valid after the `ask-chatgpt` action has been executed.

```
package action

run[res] {
    input.called[_].id == "ask-chatgpt"
    res := {
        "id": "notify-slack",
        "uses": "slack.post",
        "args": {
            "secret_webhook_url": input.env.SLACK_INCOMING_WEBHOOK,
        },
    }
}
```

### `init` rule

The `init` rule is called only once per alert. This rule is mainly used to filter out alerts by an global attribute.

#### Input

- `input.alert`: [Alert](#alert)
- `input.env`: Map of (string, string): Map of environment variables of the AlertChain process.

#### Output

- `attrs` (array, optional): Array of [Attribute](#Attribute).
- `abort` (boolean, optional): If set to true, the alert will be ignored.

### `run` rule

#### Input

An Action Policy accepts the following input:

- `input.alert`: [Alert](#alert)
- `input.env`: Map of (string, string): Map of environment variables of the AlertChain process.
- `input.seq` (number): Sequence number of actions, starting from 0.
- `input.called`: Array of [Action](#action): Actions that have already been called.

Using this input, the action policy can process the alert data and determine the most appropriate action to perform next, along with the necessary arguments and Attributes.

#### Output

After evaluating the action policy, if the next action is required, set the `action` field according to the schema of [Action](#action):

### `exit` rule

#### Input

- `input.alert`: [Alert](#alert)
- `input.seq` (number): Sequence number of actions, starting from 0.
- `input.called`: Array of [Action](#action): Actions that have already been called.
- `input.action`: [Action](#action). The previously executed action is stored here. The response from that action is included in `input.action.result`.

#### Output

- `attrs`: Array of [Attribute](#Attribute).

## Basic data structures

### Alert

- `title` (string, required): Title of the alert
- `description` (string, optional): Human-readable explanation about the alert
- `source` (string, optional): Data source
- `attrs` (array, optional): Array of [Attribute](#Attribute)
- `refs` (array, optional): Array of [Reference](#Reference)
- `namespace` (string, optional): Namespace of Attributes (attrs). Global attributes are shared among alerts and actions that have the same namespace. If not set, global attribute feature is not enabled.
- `data` (any): Original data of the alert
- `raw` (string): Pretty-printed JSON string of the alert data

### Attribute

- `key` (string, required): Name of the Attribute
- `value` (any, required): Value of the Attribute
- `id` (string, optional): ID of the Attribute. If not set, it will be assigned automatically.
- `type` (string, optional): Type of the Attribute
- `global`: (boolean, optional): If set to true, the Attribute will be available to all alert and action that has same namespace. If set to false, the Attribute will only be available to the action that created it. Default is false.
- `ttl` (number, optional): Retention period of the Attribute in seconds. It's available only when `global` is true. Default is 86400.

Within a single alert, the `key` of an Attribute can be duplicated, but the `id` must be unique. If duplicate `id`s are provided, the Attribute specified later will overwrite the earlier one. Please note that the execution order of actions within the same sequence is not guaranteed, so caution is advised when specifying IDs to avoid duplication. If you need to modify an Attribute, you can intentionally overwrite it by specifying its ID.

### Reference

Referring to an external document or service resource via URL.

- `url`: (string, required)
- `title`: (string, optional)

### Action

- `id` (string, optional): A unique ID for the action within the alert. If not specified, it will be assigned automatically. An ID should only be executed once, so do not specify an ID for actions that need to be executed multiple times. Conversely, by explicitly specifying an `id`, you can prevent an action from being executed multiple times.
- `uses` (string, required): Specify the name of the action to be launched.
- `args`: Specify the arguments for each action in a key-value format.
- `result`: When called in the `exit` rule, the result of the action is stored.
- `force`: (boolean, optional): If set to true, the workflow will be continued even if the action get error. Default is false.

NOTE: Arguments with the `secret_` prefix in `args` have a special meaning. This indicates that the value is confidential (e.g., API keys) and will not be output in logs or similar records.

## Global Attribute

The Global Attribute is a mechanism designed to share states between different alerts. It requires an external database to function. The databases presently supported are as follows:

- [Google Cloud Firestore](https://cloud.google.com/firestore)

### Use cases

The Global Attribute is designed for use in the following scenarios:

- To execute specific actions only when an alert occurs a specified number of times within a certain period. For example, if there are more than 10 accesses from the same IP address, it may be considered suspicious behavior and reported as an alert. If the number of occurrences increases further, actions such as blocking may be carried out.
- To execute different actions based on the actions taken for past alerts. For example, if similar alerts occur multiple times in a short period, notifications can be suppressed or consolidated to the same destination. This can be used to aggregate multiple alerts into a single ticket in a ticket management service, or to write messages within a thread rather than splitting them in a chat service, etc.

### Working Principle

Global Attributes are stored in an external database. Every time an alert is processed (Action Policy execution), AlertChain fetches the Global Attributes associated with the alert from the external database and saves any added or updated Global Attributes back to the database after processing.

Global Attributes are stored under a **namespace**. Alerts and actions within the same namespace can reference the same Global Attribute. The namespace is specified in the alert's `namespace` field. If no namespace is specified, Global Attributes cannot be used, and all related rules are ignored.

- Global Attributes are uniquely identified by a combination of a namespace and an Attribute.ID. If you try to create a Global Attribute with the same ID in the same namespace, the value will be overwritten.
- Global Attributes can have a specified TTL (Time To Live). If no TTL is specified, the default is 24 hours. Global Attributes that exceed their TTL are deleted.
- When values are overwritten, the TTL is updated.
- Alerts with the same namespace are always processed in series. That is, multiple alerts with the same namespace are never processed simultaneously. This ensures that Global Attributes are updated without conflict. However, the execution order of processes whose timing clashes is not guaranteed. The process that can acquire the lock the fastest will be executed first.

### Examples

First, you need to specify a namespace in the Alert Policy. The namespace is specified in the alert's `namespace` field.

```rego
package alert.my_alert

alert[msg] {
    input.name == "suspicious_action"
    msg := {
        "title": "detected suspicious action",
        "namespace": "my_namespace",
    }
}
```

The namespace can be a fixed value, or it can be dynamically determined according to the content of the alert. This allows you to store different values for each target you are checking.

```rego
        "namespace": input.user,
```

Next, you will write about the Global Attribute in the Action Policy. Global Attributes are loaded just before the `init` rule is evaluated and added to the Alert's Attributes.

```rego
package action

init[res] {
    input.alert.attrs[x].key == "called"
    input.alert.attrs[x].value == true

    res := {
        "abort": true,
    }
}
```

For example, if there is an Attribute with the key `called`, you can set `abort` to true if the value is greater than 0, thereby aborting the processing of the alert. This can be used in the `init` rule to prevent the same alert from being processed multiple times.

Here are examples of `run` and `exit`.

```rego
package action

run[job] {
    job := {
        "id": "notify-slack",
        "uses": "slack.post",
        "args": {
            "secret_webhook_url": input.env.SLACK_INCOMING_WEBHOOK,
        },
    }
}

exit[res] {
    input.action.id == "notify-slack"
    res := {
        "attrs": [
            {
                "key": "called",
                "value": true,
                "global": true,
                "ttl": 3600,
            }
        ]
    }
}
```

For `run`, you can write as usual, and of course, you can use the already loaded Global Attributes as needed. They can be used in decision conditions or as arguments for action execution.

In `exit`, you are writing the addition of Global Attributes. Here you are setting an Attribute with the key `called`, which is incorporated into the condition of the previous `init` rule. By specifying `true` in the `global` field, it is treated as a Global Attribute. The evaluation of `run` and `exit` is repeated, and at the end of all processing, only the Attributes with `global` set to `true` are saved to the database.

During this series of processes, the Action of the alert with the same namespace is not executed. Therefore, there will be no conflict within AlertChain.

## FAQ

### Why are two rules, init and exit, necessary?

Both init and exit are called in between actual actions. Therefore, you might wonder if only one of them would be sufficient.

The major difference between these two is that exit is called each time an action finishes, while init is called after all actions called from the previous run rule have completed. If you want to decide on attributes to retain based on the results of multiple actions, or decide whether to continue processing, you use init. On the other hand, if you want to update attributes based on the results of an action, or interrupt processing, you use exit.

Also, init is the first rule evaluated in the Action Policy. Therefore, it can also be used to decide whether to continue or interrupt processing by referring to information associated with past alerts (Global Attribute).
