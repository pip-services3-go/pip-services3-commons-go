package data

/*
An object that contains string translations for multiple languages.
Language keys use two-letter codes like: 'en', 'sp', 'de', 'ru', 'fr', 'pr'.
When translation for specified language does not exists it defaults to English ('en').
When English does not exists it falls back to the first defined language.

Example:
 values := MultiString.fromTuples(
     "en", "Hello World!",
     "ru", "Привет мир!"
 );

 value1 := values.get('ru'); // Result: "Привет мир!"
 value2 := values.get('pt'); // Result: "Hello World!"
*/
type MultiString map[string]string
