  if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
    var link = document.createElement('link');
    link.rel = 'stylesheet';
    link.href = '/static/css/sakura-dark.min.css';
    document.head.appendChild(link);
  } else {
    var link = document.createElement('link');
    link.rel = 'stylesheet';
    link.href = '/static/css/sakura-light.min.css';
    document.head.appendChild(link);
  }
