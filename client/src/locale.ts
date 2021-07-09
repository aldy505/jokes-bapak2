/**
 * TODO: Check user locale, then determines whether they should go to english route or indonesian route.
 */

const getLanguage = () => navigator?.languages[0] || navigator?.language || 'en';

export {}