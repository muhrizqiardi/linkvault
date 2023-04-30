import { Button, buttonVariants } from './button';

export const ButtonStories = () => (
  <>
    <div className="mb-4 flex flex-wrap gap-4">
      <Button>Variant: Primary</Button>
      <Button variant="secondary">Variant: Secondary</Button>
      <Button variant="outline">Variant: Outline</Button>
      <Button variant="destructive">Variant: Destructive</Button>
      <Button variant="ghost">Variant: Ghost</Button>
      <Button variant="link">Variant: Link</Button>
    </div>
    <div className="mb-4">
      <a href="#" className={buttonVariants({ variant: 'link' })}>
        Link with Button styling
      </a>
    </div>
  </>
);
