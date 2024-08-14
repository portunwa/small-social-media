import type { Metadata } from "next";
import { Inter } from "next/font/google";
import Navbar from '@/components/Navbar'
import { AuthProvider } from '@/context/AuthContext';
import "./globals.css";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Small Social Media App",
  description: "Next.js and Go learning project",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <AuthProvider>
        <body className={inter.className}>
          <Navbar />
          <div>{children}</div>
        </body>
      </AuthProvider>
    </html>
  );
}
