
def buildQueryString(queryParams=None):
    if queryParams is None:
        return ""

    qs = "?"
    for key, elem in queryParams:
        qs += key + "=" + str(elem) + "&"

    return qs[:-1]
