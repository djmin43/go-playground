import { rest } from 'msw';

const handlers = [
  rest.get('https://picsum.photos/v2/list', (req, res, ctx) => res(
    ctx.status(200),
    ctx.json({
      data: [
        {
          id: 1,
          author: 'tester1',
          download_url: 'tester_url1',
          
        },
        {
          id: 2,
          author: 'tester2',
          download_url: 'tester_url2',
        },
      ],
    }),
  )),
];

export { handlers }