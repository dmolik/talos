module.exports = {
  defaultLanguage: `en`,
  availableRoutesFile: `${__dirname}/routes.json`,
  contentInputFolder: `${__dirname}/content/`,
  menuInputFile: `${__dirname}/content/{version}.{lang}.json`,
  menuOutputFile: `${__dirname}/static/{version}.menu.{lang}.json`,
  sectionsOutputFile: `${__dirname}/static/{version}.sections.{lang}.json`,
  redirectsOutputFile: `${__dirname}/static/_redirects`
}
