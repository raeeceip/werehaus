import React, { useState } from 'react';
import { Button } from './button';
import { fetchInventoryReport, fetchIssueReport } from '../../src/lib/api';

const Reports = () => {
  const [reportData, setReportData] = useState(null);
  const [loading, setLoading] = useState(false);

  const loadReport = async (reportType) => {
    setLoading(true);
    try {
      const data = await (reportType === 'inventory' ? fetchInventoryReport() : fetchIssueReport());
      setReportData(data);
    } catch (error) {
      alert(`Error loading report: ${error.message}`);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="space-y-4">
      <div className="flex space-x-4">
        <Button onClick={() => loadReport('inventory')} disabled={loading}>
          Inventory Report
        </Button>
        <Button onClick={() => loadReport('issues')} disabled={loading}>
          Issue Report
        </Button>
      </div>
      {loading && <p>Loading report...</p>}
      {reportData && (
        <div className="bg-secondary p-4 rounded-md">
          <pre className="whitespace-pre-wrap">{JSON.stringify(reportData, null, 2)}</pre>
        </div>
      )}
    </div>
  );
};

export default Reports;