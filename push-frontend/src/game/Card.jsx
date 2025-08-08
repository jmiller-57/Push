export default function Card({ card, onClick }) {
  const { rank, suit } = card;

  return (
    <button className="card" onClick={onClick} title={`${rank}${suit}`}>
      <span className="rank">{rank}</span>
      <span className="suit">{suit}</span>
    </button>
  );
}