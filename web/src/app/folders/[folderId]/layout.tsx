import { SidebarBase } from '@/components/sidebar-base';

export default function FolderPageLayout(props: { children: React.ReactNode }) {
  return (
    <div className="grid lg:grid-cols-[256px_576px_minmax(0,1fr)]">
      <div className="">
        <SidebarBase className="border-r hidden lg:block" />
      </div>
      <div className="border-r">{props.children}</div>
      <div className=""></div>
    </div>
  );
}
