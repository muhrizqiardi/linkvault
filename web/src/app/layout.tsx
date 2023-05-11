import { SidebarBase } from '@/components/sidebar-base';
import '@/styles/globals.css';
import { Inter } from 'next/font/google';

const inter = Inter({ subsets: ['latin'], variable: '--font-sans' });

export default function RootLayout({
  children,
  linkDetailPanel,
}: {
  children: React.ReactNode;
  linkDetailPanel: React.ReactNode;
}) {
  return (
    <html lang="en" className={`${inter.className}`}>
      <body>
        <div className="grid lg:grid-cols-[256px_576px_minmax(0,1fr)]">
          <div className="h-screen sticky top-o">
            <SidebarBase className="border-r hidden lg:block" />
          </div>
          <div className="border-r">{children}</div>
          <div className="h-full">{linkDetailPanel}</div>
        </div>
      </body>
    </html>
  );
}
