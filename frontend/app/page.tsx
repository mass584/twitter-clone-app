"use client";

import * as React from "react";
import styles from "./page.module.css";
import HeaderContainer from "@/components/Header";
import TimelineContainer from "@/components/Timeline";
import PostTweetContainer from "@/components/PostTweet";

export default function Home() {
  return (
    <>
      <HeaderContainer />
      <main className={styles.main}>
        <PostTweetContainer />
        <TimelineContainer />
      </main>
    </>
  );
}
