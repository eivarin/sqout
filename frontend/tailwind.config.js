/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {}
	},
  daisyui: {
    themes: [
      {
        mytheme: {
          primary: '#55dde0',
          secondary: '#33658a',
          accent: '#f26419',
          neutral: '#112621',
          'base-100': '#2f4858',
          info: '#008199',
          success: '#5fe100',
          warning: '#ffc000',
          error: '#ff004a'
        }
      }
    ]
  },
	plugins: [require('daisyui')]
};
