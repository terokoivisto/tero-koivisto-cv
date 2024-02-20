# My CV

## Why would you do this?
I didn't want to spend more time updating my old CV using Photoshop, 
as it is annoying, and boring. I wanted to create a visually appealing 
CV that stands out, and I feel like this is a fitting choice.

I decided to create the frontend using Svelte, as I really like it.
It is simple, easy to use, forces to have one component per file and 
in my opinion it is much easier to read and comprehend what is happening,
than when using React.

For the backend, I wanted to try Go, as I have never used it before, 
and it seemed like an easy to pick up, which it totally turned out to be.
Definitely going to use it more.

For the database, I decided to use DynamoDB, as I don't get to use 
it too often, and I wanted to see how the AWS Go SDKs compare to 
other languages. I didn't spend hours after hours trying to figure 
out what simple things so, so it seems pretty nice compared to 
e.g. Boto3.

## With what?
**Frontend**: [Svelte](https://svelte.dev/) \
**Backend**: [go](https://go.dev/) \
**Database**: [DynamoDB (locally)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.DownloadingAndRunning.html)

## How can I drive this thing?

##### Prerequisites
- Docker
- Node
- pnpm
- Go

You can run the backend and DynamoDB in Docker with 
```bash
$ cd backend
$ cp .env.sample .env
$ docker-compose up --build
```

You can run the frontend in dev mode with
```bash
$ cd frontend
$ pnpm install
$ pnpm run dev
```

You can then navigate to the URL given by Vite, usually http://localhost:5173
