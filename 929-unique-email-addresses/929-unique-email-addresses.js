/**
 * @param {string[]} emails
 * @return {number}
 */
var numUniqueEmails = function(emails) {
  return new Set(emails.map(email => {
    const [local, domain] = email.split('@')
    const [part, _] = local.split('+')
    return `${part.split('.').join('')}@${domain}`
  })).size
};