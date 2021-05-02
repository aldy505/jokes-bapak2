function socialTags({
  title, description, url, image, keywords, author,
}) {
  const meta = [];
  if (title) {
    meta.push({ property: 'og:title', content: title });
    meta.push({ name: 'twitter:title', content: title });
  }
  if (description) {
    meta.push({ property: 'og:description', content: description });
    meta.push({ name: 'twitter:description', content: description });
  }
  if (url) {
    meta.push({ property: 'og:url', content: url });
    meta.push({ name: 'twitter:url', content: url });
  }
  if (image) {
    meta.push({ property: 'og:image', content: image });
    meta.push({ name: 'twitter:image', content: image });
  }
  if (keywords) {
    meta.push({ name: 'keywords', content: keywords });
  }
  if (author) {
    meta.push({ property: 'og:author', content: author });
    meta.push({ name: 'twitter:author', content: author });
  }
  return meta;
}

function injectToHead(document, headItems) {
  const { head } = document;
  for (let i = 0; i < headItems.length; i += 1) {
    const meta = document.createElement('meta');
    meta.push(headItems[i]);
    head.appendChild(meta);
  }
}

export default { injectToHead, socialTags };
