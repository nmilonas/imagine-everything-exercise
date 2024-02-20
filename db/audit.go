package db

// BatchInsertAuditDataInTransaction will create a db transaction in order to store the audit data. Since large amounts of data
// could be added at once, batching will be used to reduce the overall usage of the db, and the transaction will serve as a safety net
// in case there is an issue with certain batch additons.
func BatchInsertAuditDataInTransaction() {}

// GetAuditRecordsByTimeRangeAndEmail will fetch audit records for a specfic email and a given time range.
func GetAuditRecordsByTimeRangeAndEmail() {}
