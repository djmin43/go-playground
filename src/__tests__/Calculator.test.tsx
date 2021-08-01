import { render, unmountComponentAtNode } from "react-dom";
import { act } from "react-dom/test-utils";

import Calculator from '../components/Caclulator'

let container: any = null;
beforeEach(() => {
  // setup a DOM element as a render target
  container = document.createElement("div");
  document.body.appendChild(container);
});

afterEach(() => {
  // cleanup on exiting
  unmountComponentAtNode(container);
  container.remove();
  container = null;
});

it('redners a title', () => {
  act(() => {
    render(<Calculator />, container)
  })
  expect(container.textContent).toBe('Calculator')
})