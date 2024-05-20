"use client";

import { configuredAxios } from "@/lib/handler";
import * as React from "react";

export const Login = (props: {
  email: string;
  isError: boolean;
  onChange: (v: string) => void;
  onSubmit: () => void;
}) => (
  <>
    <p style={{ marginBottom: "1rem", textAlign: "center" }}>
      メールアドレスでログインしてください
      <br />
      今はまだサインアップ機能がないので、seedデータのメールアドレスを指定してね(sorry)
    </p>
    <div style={{ display: "flex", marginBottom: "1rem" }}>
      <input
        style={{ marginRight: "0.25rem" }}
        onChange={(e) => {
          const v = e.currentTarget.value;
          props.onChange(v);
        }}
        value={props.email}
      />
      <button onClick={props.onSubmit}>ログイン</button>
    </div>
    {props.isError && <p>ユーザーが見つかりません</p>}
  </>
);

const LoginContainer = () => {
  const [email, setEmail] = React.useState<string>("");
  const [isError, setIsError] = React.useState<boolean>(false);

  const onChange = React.useCallback((v: string) => setEmail(v), []);

  const onSubmit = React.useCallback(() => {
    (async () => {
      try {
        setIsError(false);
        const requestBody = { email: email };
        await configuredAxios.post(`/login`, requestBody);
        // ログインできたらホームに飛ばす
        location.href = "/";
      } catch (e: any) {
        if (e.response.status === 401) {
          setIsError(true);
        } else {
          // システムエラーのハンドリング考える
          console.log(e);
        }
      }
    })();
  }, [email]);

  return (
    <Login
      email={email}
      isError={isError}
      onChange={onChange}
      onSubmit={onSubmit}
    />
  );
};

export default LoginContainer;
