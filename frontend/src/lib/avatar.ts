const COLORS = [
	'#1da1f2',
	'#e0245e',
	'#17bf63',
	'#f45d22',
	'#794bc4',
	'#0d8ecf',
	'#aa8ed6',
	'#f5a623'
];

export function getInitials(name: string): string {
	const trimmed = name.trim();
	if (!trimmed) return '?';
	const parts = trimmed.split(/\s+/);
	if (parts.length === 1) return parts[0][0].toUpperCase();
	return (parts[0][0] + parts[parts.length - 1][0]).toUpperCase();
}

export function getAvatarColor(seed: string): string {
	let hash = 0;
	for (let i = 0; i < seed.length; i++) {
		hash = (hash << 5) - hash + seed.charCodeAt(i);
		hash |= 0;
	}
	return COLORS[Math.abs(hash) % COLORS.length];
}
