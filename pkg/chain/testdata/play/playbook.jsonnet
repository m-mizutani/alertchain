{
  scenarios: [
    {
      id: 's1',
      title: 'Scenario 1',
      alert: {
        class: 'threat',
      },
      schema: 'my_test',
      results: {
        my_action: [
          {
            index: 'first',
          },
          {
            index: 'second',
          },
        ],
      },
    },
  ],
}
