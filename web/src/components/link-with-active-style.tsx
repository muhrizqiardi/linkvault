'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';

interface LinkWithActiveStyleProps
  extends React.AnchorHTMLAttributes<HTMLAnchorElement> {
  href: string;
  activeClassName: string;
  className: string;
}

export function LinkWithActiveStyle({
  activeClassName,
  className,
  href,
  ...restOfProps
}: LinkWithActiveStyleProps) {
  const pathname = usePathname();

  const activeStyledLink = (
    <Link href={href} legacyBehavior>
      <a className={activeClassName} {...restOfProps} />
    </Link>
  );
  const inactiveStyledLink = (
    <Link href={href} legacyBehavior>
      <a className={className} {...restOfProps} />
    </Link>
  );

  if (pathname === null) return inactiveStyledLink;
  if (pathname === href) return activeStyledLink;

  return inactiveStyledLink;
}
