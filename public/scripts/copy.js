function copyToClipboard() {
    const timestamp = document.getElementById('utc-timestamp').textContent;
    navigator.clipboard.writeText(timestamp);
}

function updateTimes() {
    fetch('/api/times')
        .then(response => response.json())
        .then(data => {
            document.getElementById('utc-timestamp').textContent = data.utc;
            // Update other time elements
            Object.entries(data.times).forEach(([zone, time]) => {
                const element = document.getElementById(`time-${zone}`);
                if (element) element.textContent = time;
            });
        });
}