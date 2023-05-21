export function formatDate(date : Date) {
    if (!isValid(date)) {
        return undefined
    }
    const formatter = Intl.DateTimeFormat("en-GB", {
        day: "2-digit",
        month: "2-digit",
        year: "numeric"
    })
    return formatter.format(date);
}

export function formatDateTime(dateTime : Date) {
    if (!isValid(dateTime)) {
        return undefined
    }
    const formatter = Intl.DateTimeFormat("en-GB", {
        day: "2-digit",
        month: "2-digit",
        year: "numeric",
        hour: "2-digit",
        minute: "2-digit",
        hour12: false
    })
    return formatter.format(dateTime);
}

function isValid(date: Date) {
    if (Object.prototype.toString.call(date) !== "[object Date]") {
        return false
    }
    if (isNaN(date.getTime())) {
        return false
    } else {
        return true
    }
}

function addHours() {
    
}