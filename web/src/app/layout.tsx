import { SidebarBase } from '@/components/sidebar-base';
import '@/styles/globals.css';
import { Inter } from 'next/font/google';

const inter = Inter({ subsets: ['latin'] });

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en" className={`${inter.className}`}>
      <body>
        <div className="grid lg:grid-cols-[256px_576px_minmax(0,1fr)]">
          <div className="">
            <SidebarBase className="border-r hidden lg:block" />
          </div>
          <div className="border-r">{children}</div>
          <div className=""></div>
        </div>
      </body>
    </html>
  );
}
