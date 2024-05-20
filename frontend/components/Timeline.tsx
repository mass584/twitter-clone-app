"use client";

import { formatToTimeZone, parseFromTimeZone } from "date-fns-timezone";
import { configuredAxios } from "@/lib/handler";
import * as React from "react";

const TIMEZONE = "Asia/Tokyo";

const PER = 10;

type Tweet = {
  id: number;
  displayName: number;
  textContents: string;
  tweetAt: string;
};

const formattedDatetime = (date_string: string) => {
  const date = parseFromTimeZone(date_string, { timeZone: TIMEZONE });

  return formatToTimeZone(date, "YYYY-MM-DD HH:mm:ss", { timeZone: TIMEZONE });
};

export const Timeline = (props: {
  displayedTweets: ReadonlyArray<Tweet>;
  onLoadLatestTimeline: () => void;
  onLoadOldTimeline: () => void;
}) => (
  <>
    <button
      onClick={props.onLoadLatestTimeline}
      style={{ marginBottom: "1rem" }}
    >
      最新のツイートを読み込む
    </button>
    {props.displayedTweets.map((displayedTweet, index) => (
      <div key={index} style={{ marginBottom: "1rem", width: "100%" }}>
        <hr />
        <p style={{ fontSize: "0.75rem" }}>{`${
          displayedTweet.displayName
        }（${formattedDatetime(displayedTweet.tweetAt)}）`}</p>
        {displayedTweet.textContents.split("\n").map((line, index) => (
          <p key={index}>{line}</p>
        ))}
      </div>
    ))}
    <button onClick={props.onLoadOldTimeline}>過去のツイートを読み込む</button>
  </>
);

const TimelineContainer = () => {
  const [displayedTweets, setDisplayedTweets] = React.useState<
    ReadonlyArray<Tweet>
  >([]);
  const [page, setPage] = React.useState<number>(1);

  // 過去のタイムラインページを取得してくるためのhook
  React.useEffect(() => {
    (async () => {
      try {
        // 今のapi実装だと、フォロワーが新たにツイートするとにオフセット位置がずれてしまう
        // オフセット基準を日時で固定できるようにして、冪統性のあるapiに拡張するべき
        const { data } = await configuredAxios.get(
          `/timeline?per=${PER}&page=${page}`
        );

        const fetchedTweets: ReadonlyArray<Tweet> = data.timeline.Tweets.map(
          (fetchedTweet: any) => ({
            id: fetchedTweet.Id,
            displayName: fetchedTweet.User.DisplayName,
            textContents: fetchedTweet.TextContents,
            tweetAt: fetchedTweet.TweetAt,
          })
        );

        setDisplayedTweets((pre) => {
          return [
            ...pre,
            ...fetchedTweets.filter(
              // 同じ内容のツイートを複数個表示しないようにフィルタリング
              // apiを冪等にできればこの処理は不要になる
              (fetchedTweet: any) =>
                !pre.some(
                  (displayedTweet) => displayedTweet.id === fetchedTweet.id
                )
            ),
          ];
        });
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
  }, [page]);

  // 雑にページをリロードしているが、過去タイムラインを捨ててしまうためよくない。
  // ロード済みのツイートをキャッシュしたまま最新のものだけロードする仕組みを考える。
  const onLoadLatestTimeline = React.useCallback(() => location.reload(), []);
  const onLoadOldTimeline = React.useCallback(
    () => setPage((pre) => pre + 1),
    []
  );

  return (
    <Timeline
      displayedTweets={displayedTweets}
      onLoadLatestTimeline={onLoadLatestTimeline}
      onLoadOldTimeline={onLoadOldTimeline}
    />
  );
};

export default TimelineContainer;
