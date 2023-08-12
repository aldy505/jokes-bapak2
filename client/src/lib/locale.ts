import { addMessages, getLocaleFromNavigator, getLocaleFromQueryString, init } from 'svelte-i18n';

import en from '../languages/en.json';
import id from '../languages/id.json';

addMessages('en', en);
addMessages('en-US', en);
addMessages('en-GB', en);
addMessages('id', id);
addMessages('id-ID', id);

init({
  fallbackLocale: 'en',
  initialLocale: getLocaleFromQueryString('lang') || getLocaleFromNavigator(),
});
