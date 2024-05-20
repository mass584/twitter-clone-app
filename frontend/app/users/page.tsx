"use client";

import * as React from "react";
import styles from "../page.module.css";
import HeaderContainer from "@/components/Header";
import UsersContainer from "@/components/Users";

export default function Home() {
  return (
    <>
      <HeaderContainer />
      <main className={styles.main}>
        <UsersContainer />
      </main>
    </>
  );
}
