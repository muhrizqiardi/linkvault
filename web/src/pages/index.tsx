import { SidebarBase } from '@/components/sidebar-base';
import { TopNavBase } from '@/components/top-nav-base';

export default function Home() {
  return (
    <>
      <div className="grid grid-cols-[256px_576px_minmax(0,1fr)]">
        <div className="">
          <SidebarBase className="border-r" />
        </div>
        <div className="border-r">
          <TopNavBase />
        </div>
        <div className=""></div>
      </div>
    </>
  );
}
