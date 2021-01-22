
let redirectPage = "https://login.microsoftonline.com/common/oauth2/v2.0/authorize?"

function addToRedirectPage(ParaName: string, ParaDetails: string) {
    redirectPage = redirectPage + "&" + ParaName + "=" + ParaDetails
};

function addToRedirectPageFirst(ParaName: string, ParaDetails: string) {
    redirectPage = redirectPage + ParaName + "=" + ParaDetails
};

addToRedirectPageFirst("client_id", "a2933e57-93e2-4991-b1b0-d8dcf4c7a2a9")
addToRedirectPage("redirect_uri", "http://localhost:9000/mainPage")
addToRedirectPage("scope", "Mail.Read offline_access")
addToRedirectPage("response_type", "code")
addToRedirectPage("response_mode", "query")

function oauthRequest() {
    window.location.href = redirectPage
};


export { oauthRequest }