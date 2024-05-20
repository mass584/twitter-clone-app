/** @type {import('next').NextConfig} */
const nextConfig = {
  env: {
    API_SERVER_HOST: process.env.API_SERVER_HOST ?? "http://localhost:8080",
  },
};

export default nextConfig;
