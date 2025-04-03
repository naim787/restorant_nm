<script>
	import '@carbon/charts-svelte/styles.css';
	import { LineChart } from '@carbon/charts-svelte';

	// Data dengan Tanggal & Waktu
	let data = [
		{ group: 'Dataset 1', date: '2024-01-01T08:00:00', value: 10000 }, // 08:00
		{ group: 'Dataset 1', date: '2024-01-01T10:30:00', value: 30000 }, // 10:30
		{ group: 'Dataset 1', date: '2024-01-01T13:15:00', value: 50000 }, // 13:15
		{ group: 'Dataset 1', date: '2024-01-01T15:45:00', value: 25000 }, // 15:45
		{ group: 'Dataset 1', date: '2024-01-01T18:00:00', value: 45000 }  // 18:00
	];

    setInterval(() => {
        let now = new Date();
		let newData = [];
		for (let i = 9; i >= 0; i--) {
			let time = new Date(now.getTime() - i * 2 * 60000);
			newData.push({
				group: 'Dataset 1',
				date: time.toISOString(),
				value: Math.floor(Math.random() * 50000) + 10000
			});
		}
        console.log(newData)
    }, 2000)

	let options = {
		theme: 'g90', // Tema gelap (bisa pakai 'white' untuk terang)
		title: 'Line Chart dengan Tanggal & Waktu',
		height: '400px',
		axes: {
			left: { mapsTo: 'value', title: 'Jumlah' },
			bottom: { mapsTo: 'date', scaleType: 'time', title: 'Waktu' }
		},
		curve: 'curveMonotoneX', // Garis lebih halus
		containerBackgroundStyle: { fill: 'transparent' }, // Background transparan
		colorScale: { colors: ['#ffffff'] } // Warna garis putih (cocok untuk tema gelap)
	};
</script>

<LineChart {data} {options} />
