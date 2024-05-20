"use client";

export const Header = () => (
  <header
    style={{
      display: "flex",
      flexDirection: "row",
      alignItems: "center",
      padding: "0 6rem",
      minHeight: "6rem",
    }}
  >
    <a
      href="/"
      style={{
        padding: "0 1rem",
      }}
    >
      HOME
    </a>
    <a
      href="/users"
      style={{
        padding: "0 1rem",
      }}
    >
      USER
    </a>
    <a
      href="/login"
      style={{
        marginLeft: "auto",
        padding: "0 1rem",
      }}
    >
      LOGIN
    </a>
  </header>
);

const HeaderContainer = Header;

export default HeaderContainer;
