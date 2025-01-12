let text = Array.from(document.querySelector('.entry-content').querySelectorAll('p'))
    .map(p => p.textContent.trim())
    .join('\n');
console.log(text);