"use client";

import * as React from "react";
import styles from "../page.module.css";
import HeaderContainer from "@/components/Header";
import LoginContainer from "@/components/Login";

export default function Home() {
  return (
    <>
      <HeaderContainer />
      <main className={styles.main}>
        <LoginContainer />
      </main>
    </>
  );
}
