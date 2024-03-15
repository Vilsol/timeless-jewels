import { addMessages, getLocaleFromNavigator, init } from 'svelte-i18n';

import locales from './locales.json';

addMessages('en', locales.en);
addMessages('zh', locales.zh);

init({
  fallbackLocale: 'en',
  initialLocale: getLocaleFromNavigator()
});
