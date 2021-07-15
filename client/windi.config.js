import { defineConfig } from 'windicss/helpers';

export default defineConfig({
  darkMode: 'media',
  theme: {
    extend: {
      fontFamily: {
        code: ['"Fira Mono"', '"Source Code Pro"', '"Lucida Console"', '"Courier New"', 'Courier', 'sans-serif'],
        body: ['Rubik', 'Asap', 'Barlow', 'Arial', 'sans-serif'],
      },
      flex: {
        2: '2 2 0%',
        3: '3 3 0%',
        4: '4 4 0%',
        5: '5 5 0%',
      },
      colors: {
        chetwode: {
          50: '#f9fafe',
          100: '#f3f4fe',
          200: '#e1e4fc',
          300: '#cfd4f9',
          400: '#abb4f5',
          500: '#8794f1',
          600: '#7a85d9',
          700: '#656fb5',
          800: '#515991',
          900: '#424976',
        },
        dodger: {
          50: '#f4f9fe',
          100: '#e8f2fd',
          200: '#c6dffa',
          300: '#a4cbf7',
          400: '#5fa5f2',
          500: '#1b7eec',
          600: '#1871d4',
          700: '#145fb1',
          800: '#104c8e',
          900: '#0d3e74',
        },
        lavender: {
          50: '#fefbff',
          100: '#fef6fe',
          200: '#fceafd',
          300: '#fbddfb',
          400: '#f7c3f8',
          500: '#f4a9f5',
          600: '#dc98dd',
          700: '#b77fb8',
          800: '#926593',
          900: '#785378',
        },
      },
    },
  },
});
