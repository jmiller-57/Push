export default function DrawPile({ count }) {
  return (
    <div className="draw">
      <div className="title">Draw Pile</div>
      <div className="stack">{count ?? 0}</div>
    </div>
  );
}