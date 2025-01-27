let text = Array.from(document.querySelector('.entry-content').querySelectorAll('p'))
    .map(p => p.textContent.trim())
    .join('\n')
    .replace(/[\[\]]/g, '');
console.log(text);
