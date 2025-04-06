function fetchSuggestions(input) {
    if (input.length < 2) {
        document.getElementById('suggestions').innerHTML = '';
        return;
    }

    fetch(`/suggest?prefix=${input}`)
        .then(response => response.json())
        .then(suggestions => {
            const list = document.getElementById('suggestions');
            list.innerHTML = suggestions
                .map(name => `<li onclick="selectSuggestion('${name}')">${name}</li>`)
                .join('');
        });
}

function selectSuggestion(name) {
    document.getElementById('emojiInput').value = name;
    document.getElementById('suggestions').innerHTML = '';
}

function submitForm() {
    const input = document.getElementById('emojiInput').value;

    // Simulate form submission (e.g., send data to server via fetch)
    fetch('/', {
        method: 'POST',
        headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
        body: `input-text=${encodeURIComponent(input)}`
    })
        .then(response => response.text())
        .then(data => {
            console.log('Form submitted successfully:', data);
            // Optionally update UI or retain input value
        })
        .catch(error => console.error('Error submitting form:', error));
}

function addEmoji() {
    const name = document.getElementById('newEmojiName').value;
    const symbol = document.getElementById('newEmojiSymbol').value;

    fetch('/api/v1/emojis', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name, symbol })
    })
        .then(response => {
            if (!response.ok) {
                return response.text().then(text => { throw new Error(text) });
            }
            return response.json();
        })
        .then(data => {
            document.getElementById('addEmojiStatus').innerHTML = 
                `<div class="success">${data.message}</div>`;
        })
        .catch(error => {
            document.getElementById('addEmojiStatus').innerHTML = 
                `<div class="error">${error.message}</div>`;
        });
}

function clearInput() {
    document.getElementById('emojiInput').value = '';
    document.getElementById('suggestions').innerHTML = '';
    document.getElementById('output-container').innerHTML = '';
}
function clearAddEmojiForm() {
    document.getElementById('newEmojiName').value = '';
    document.getElementById('newEmojiSymbol').value = '';
    document.getElementById('addEmojiStatus').innerHTML = '';
}