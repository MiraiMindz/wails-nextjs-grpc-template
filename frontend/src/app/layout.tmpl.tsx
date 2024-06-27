import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";

import { ThemeProvider } from "@/components/ThemeProvider";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
    title: "{{.ProjectName}}",
    description: "{{.ProjectName}}",
};

export default function RootLayout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <html lang="en" suppressHydrationWarning>
            <body className={`${inter.className} w-screen h-screen scroll-smooth`}>
                <ThemeProvider
                    attribute="class"
                    defaultTheme="system"
                    enableSystem
                >
                    <div className="w-full h-full bg-gradient-to-tr from-slate-50 to-slate-300 dark:from-slate-900 dark:to-slate-700">
                        {children}
                    </div>
                </ThemeProvider>
            </body>
        </html>
    );
}
