<template>
  <div id="coming-soon">
    <img id="logo" alt="tricolor bull logo" src="./assets/logo.svg"/>
    <h1>a0b.io</h1>
    <h2>Software engineering you can trust.</h2>
    <p>Coming Soon!</p>
    <p>Check out owner/operator Aaron Browne's <a href="https://www.linkedin.com/in/aaron0browne/">LinkedIn</a>.</p>
  </div>
</template>

<style>
html {
  background: linear-gradient(
      rgba(0, 0, 0, 0.7),
      rgba(0, 0, 0, 0.7)
    ),
    url(assets/phl-skyline.jpg) no-repeat center center fixed;
  -webkit-background-size: cover;
  -moz-background-size: cover;
  -o-background-size: cover;
  background-size: cover;
}

#coming-soon {
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  position: absolute;
}

#logo {
  left: 50%;
  transform: translate(-50%);
  position: relative;
  width: 20rem;
}

#coming-soon * {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  text-align: center;
  color: white;
  margin-block-start: 0;
  margin-block-end: 0;
}

#coming-soon h1 {
  font-size: 4rem;
  text-align: center;
  margin-bottom: .25rem;
}

#coming-soon h2 {
  font-size: 2rem;
}

#coming-soon p:first-of-type {
  margin-top: 4rem;
}

#coming-soon p {
  font-family: 'Roboto', Helvetica, Arial, sans-serif;
  font-size: 1.5rem;
}

#coming-soon a {
  font-family: 'Roboto', Helvetica, Arial, sans-serif;
  font-size: 1.5rem;
}
</style>
