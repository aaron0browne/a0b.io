// Copyright (C) 2019 Aaron Browne
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// A copy of the license can be found in the LICENSE file and at
//         https://www.gnu.org/licenses/agpl.html

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
body {
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

body:after{
  content: "";
  position: fixed; /* stretch a fixed position to the whole screen */
  top: 0;
  height: 100vh; /* fix for mobile browser address bar appearing disappearing */
  left: 0;
  right: 0;
  z-index: -1; /* needed to keep in the background */
  background: linear-gradient(
      rgba(0, 0, 0, 0.7),
      rgba(0, 0, 0, 0.7)
    ),
    url(assets/phl-skyline.jpg) center center;
  -webkit-background-size: cover;
  -moz-background-size: cover;
  -o-background-size: cover;
  background-size: cover;
}

#coming-soon {
  margin-top: 15vh;
  left: 50%;
  transform: translate(-50%);
  position: relative;
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

#coming-soon p:last-of-type {
  margin-bottom: 4rem;
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
