"use client";

import { configuredAxios } from "@/lib/handler";
import * as React from "react";

export const PostTweet = (props: {
  editingTweet: string;
  onChangeTweet: (text: string) => void;
  onSubmit: () => void;
}) => (
  <div style={{ display: "flex", marginBottom: "1rem" }}>
    <textarea
      cols={54}
      onChange={(event) => {
        const t = event.target.value;
        props.onChangeTweet(t);
      }}
      rows={5}
      value={props.editingTweet}
    />
    <button disabled={props.editingTweet === ""} onClick={props.onSubmit}>
      ポスト
    </button>
  </div>
);

const PostTweetContainer = () => {
  const [editingTweet, setEditingTweet] = React.useState<string>("");

  const onChangeTweet = React.useCallback((text: string) => {
    const textLength = text.length;
    if (textLength < 256) {
      setEditingTweet(() => text);
    } else {
      setEditingTweet(() => text.substring(0, 255));
    }
  }, []);

  const onSubmit = React.useCallback(() => {
    (async () => {
      try {
        const requestBody = { text_contents: editingTweet };
        await configuredAxios.post(`/tweets`, requestBody);

        // 雑にページをリロードしているが、過去タイムラインを捨ててしまうためよくない。
        // ロード済みのツイートをキャッシュしたまま最新のものだけロードする仕組みを考える。
        location.reload();
      } catch (e: any) {
        if (e.response.status === 401) {
          // セッションがなければログイン画面へ
          location.href = "/login";
        } else {
          // システムエラーのハンドリング考える
          console.error(e);
        }
      }
    })();
  }, [editingTweet]);

  return (
    <PostTweet
      editingTweet={editingTweet}
      onChangeTweet={onChangeTweet}
      onSubmit={onSubmit}
    />
  );
};

export default PostTweetContainer;
