"use client";

import { configuredAxios } from "@/lib/handler";
import * as React from "react";

const PER = 2;

type FollowableUser = {
  user: {
    id: number;
    displayName: string;
  };
  isFollowed: boolean;
};

export const Users = (props: {
  followableUsers: ReadonlyArray<FollowableUser>;
  loadingUserIds: ReadonlyArray<number>;
  onToggleFollow: (userId: number, follow: boolean) => void;
}) => (
  <>
    {props.followableUsers.map((followableUser) => (
      <div
        style={{ display: "flex", padding: "1rem" }}
        key={followableUser.user.id}
      >
        <div style={{ padding: "0.5rem" }}>
          {followableUser.user.displayName}
        </div>
        <button
          disabled={props.loadingUserIds.includes(followableUser.user.id)}
          onClick={() =>
            props.onToggleFollow(
              followableUser.user.id,
              !followableUser.isFollowed
            )
          }
        >
          {followableUser.isFollowed ? "フォロー解除" : "フォローする"}
        </button>
      </div>
    ))}
  </>
);

const UsersContainer = () => {
  const [displayedFollowableUsers, setDisplayedFollowableUsers] =
    React.useState<ReadonlyArray<FollowableUser>>([]);
  const [loadingUserIds, setLoadingUserIds] = React.useState<
    ReadonlyArray<number>
  >([]);

  // フォロー可能なユーザー一覧を取得してくるためのhook
  React.useEffect(() => {
    (async () => {
      try {
        // 表示名で検索できるようにしたい
        const { data } = await configuredAxios.get(`/followable_users`);

        const fetched: ReadonlyArray<FollowableUser> =
          data.followable_users.map((followableUser: any) => ({
            user: {
              id: followableUser.User.Id,
              displayName: followableUser.User.DisplayName,
            },
            isFollowed: followableUser.IsFollowed,
          }));

        setDisplayedFollowableUsers(fetched);
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
  }, []);

  const onToggleFollow = React.useCallback(
    (userId: number, follow: boolean) => {
      (async () => {
        setLoadingUserIds((pre) => [...pre, userId]);

        if (follow) {
          const requestBody = { target_user_id: userId };
          await configuredAxios.post(`/followable_users`, requestBody);
        } else {
          await configuredAxios.delete(`/followable_users/${userId}`);
        }
        setDisplayedFollowableUsers((pre) =>
          pre.map((followableUser) =>
            followableUser.user.id === userId
              ? { ...followableUser, isFollowed: follow }
              : followableUser
          )
        );

        setLoadingUserIds((pre) => pre.filter((id) => id !== userId));
      })();
    },
    []
  );

  return (
    <Users
      followableUsers={displayedFollowableUsers}
      loadingUserIds={loadingUserIds}
      onToggleFollow={onToggleFollow}
    />
  );
};

export default UsersContainer;
