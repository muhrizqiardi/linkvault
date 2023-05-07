'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';

interface LinkWithActiveStyleProps
  extends React.AnchorHTMLAttributes<HTMLAnchorElement> {
  href: string;
  activeClassName: string;
}

export function LinkWithActiveStyle({
  activeClassName,
  ...restOfProps
}: LinkWithActiveStyleProps) {
  const pathname = usePathname();

  if (pathname === null || pathname.startsWith(restOfProps.href))
    return <Link {...restOfProps} />;

  return <Link {...restOfProps} className={activeClassName} />;
}
