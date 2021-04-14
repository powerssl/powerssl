window.addEventListener('message', function (event) {
  if (event.origin !== powerssl.webAppURI) return;
  if (typeof event.data != 'object' || event.data === null) return;
  console.log(event.data);
  switch (event.data.action) {
    case 'logout':
      $.post('/logout').done(function () {
        jwtSet(null)
      });
      break;
  }
});

$(document).ready(function() {
  const queryString = window.location.search;
  const urlParams = new URLSearchParams(queryString);
  let username = getCookie('username');
  if (username !== '' && !urlParams.has('logout')) {
    $.get('/jwt', {
      sub: username,
    }).done(jwtSet);
  }

  $('#usernameAndEmail').submit(function(event) {
    event.preventDefault();

    let username = $('#username').val();
    let password = $('#password').val();

    $('#usernameAndEmail').trigger('reset');

    $.post('/login', {
      username: username,
      password: password,
    }).done(jwtSet);
  });

  $('#github').submit(function(event) {
    event.preventDefault();

    $.post('/login', {
      provider: 'github',
    }).done(function (data) {
      if (typeof data != 'object' || data === null) return;
      switch (data.action) {
        case 'redirect':
          httpRedirect(data);
          break;
        default:
          console.error(data);
          break;
      }
    });
  });
});

let httpRedirect = function(data) {
  let message = {
    action: 'http.redirect',
    url: data.url,
  };
  postMessage(message);
};

let jwtSet =  function(data) {
  let message = {
    action: 'auth.token',
    jwt: data,
  };
  postMessage(message);
};

function getCookie(cname) {
  let name = cname + '=';
  let decodedCookie = decodeURIComponent(document.cookie);
  let ca = decodedCookie.split(';');
  for(let i = 0; i <ca.length; i++) {
    let c = ca[i];
    while (c.charAt(0) === ' ') {
      c = c.substring(1);
    }
    if (c.indexOf(name) === 0) {
      return c.substring(name.length, c.length);
    }
  }
  return '';
}

function postMessage(message) {
  try {
    window.parent.postMessage(message, powerssl.webAppURI);
  } catch (e) {
    if (e.name === 'RangeError') {
      window.location.replace(powerssl.webAppURI);
    } else {
      throw e;
    }
  }
}
