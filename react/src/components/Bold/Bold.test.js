import React from 'react'
import renderer from 'react-test-renderer'
import Bold from './Bold'

test('renders correctly', () => {
  const component = renderer.create(
    <Bold>Test</Bold>,
  )
  const tree = component.toJSON()
  expect(tree).toMatchSnapshot()
})