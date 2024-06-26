import { Style } from 'hono/css'
import { jsxRenderer } from 'hono/jsx-renderer'
import { Script } from 'honox/server'
// @ts-ignore
import styles from "../style.css?url"
import wasm from "../wasm/wasm_exec.js?url"


export default jsxRenderer(({ children, title }) => {
  return (
    <html lang="ja">
      <head>
        <meta charset="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <title>{title}</title>
        {/* {import.meta.env.PROD ? (
          <link href='static/assets/style.css' rel='stylesheet' />
        ) : (
          <link href='/app/style.css' rel='stylesheet' />
        )} */}
        <link href={styles} rel="stylesheet" />
        <link rel="icon" href="/favicon.ico" />
        <Script src={wasm} />
        <Script src="/app/client.ts" async />
        <Style />
      </head>
      <body>{children}</body>
    </html>
  )
})
