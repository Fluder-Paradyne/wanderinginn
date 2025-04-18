const chapterContent = document.querySelector('.chapter-inner.chapter-content');

let extractedText = '';

if (chapterContent) {
  const paragraphs = chapterContent.querySelectorAll('p');
  paragraphs.forEach(paragraph => {
    const spans = paragraph.querySelectorAll('span');
    spans.forEach(span => {
      extractedText += span.textContent + '\n';
    });
  });
}

console.log(extractedText);
