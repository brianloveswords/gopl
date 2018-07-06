[...document.getElementsByTagName("pre")]
  .forEach(e => e.addEventListener('blur', () => Prism.highlightAll()));
