export default function BaseLayout(props: { children: React.ReactNode }) {
  return (
    <div className="drawer drawer-mobile">
      <input id="main-drawer" type="checkbox" className="drawer-toggle" />
      <div className="drawer-content">{props.children}</div>
      <div className="drawer-side">
        <label htmlFor="main-drawer" className="drawer-overlay"></label>
        <div className="bg-base-200 text-base-content w-64">
          <ul className="menu mt-4">
            <li className="menu-title">
              <span>
                Folders
              </span>
            </li>
            <li>
              <a>All items</a>
            </li>
            <li>
              <a>Folder 1</a>
            </li>
          </ul>

          <div className="mt-4 px-4">
            <a href="#" className="badge badge-outline">
              #videos
            </a>
          </div>
        </div>
      </div>
    </div>
  );
}
