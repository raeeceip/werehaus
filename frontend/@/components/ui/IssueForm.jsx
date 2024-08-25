import React, { useState, useEffect } from 'react';
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { Label } from "../ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";
import { fetchItems, fetchLocations, requestIssue } from "../../../src/lib/api";

const IssueForm = () => {
  const [items, setItems] = useState([]);
  const [locations, setLocations] = useState([]);
  const [formData, setFormData] = useState({
    item_id: '',
    quantity: '',
    from_location_id: '',
    to_location_id: ''
  });
  const [issueResult, setIssueResult] = useState('');

  useEffect(() => {
    loadFormData();
  }, []);

  const loadFormData = async () => {
    try {
      const [itemsResponse, locationsResponse] = await Promise.all([fetchItems(), fetchLocations()]);
      setItems(itemsResponse.items);
      setLocations(locationsResponse);
    } catch (error) {
      console.error("Error loading form data:", error);
      setIssueResult('Error loading form data. Please refresh the page.');
    }
  };

  const handleChange = (name, value) => {
    setFormData(prevData => ({ ...prevData, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await requestIssue(formData);
      setIssueResult('Issue request submitted successfully');
      setFormData({
        item_id: '',
        quantity: '',
        from_location_id: '',
        to_location_id: ''
      });
    } catch (error) {
      console.error("Error requesting issue:", error);
      setIssueResult('Error submitting issue request. Please try again.');
    }
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4 max-w-md">
      <div className="space-y-2">
        <Label htmlFor="item_id">Select Item</Label>
        <Select name="item_id" onValueChange={(value) => handleChange('item_id', value)}>
          <SelectTrigger>
            <SelectValue placeholder="Select an item" />
          </SelectTrigger>
          <SelectContent>
            {items.map(item => (
              <SelectItem key={item.id} value={item.id.toString()}>
                {item.name} (Quantity: {item.quantity})
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
      </div>

      <div className="space-y-2">
        <Label htmlFor="quantity">Quantity</Label>
        <Input
          type="number"
          id="quantity"
          name="quantity"
          required
          min="1"
          value={formData.quantity}
          onChange={(e) => handleChange('quantity', e.target.value)}
        />
      </div>

      <div className="space-y-2">
        <Label htmlFor="from_location_id">From Location</Label>
        <Select name="from_location_id" onValueChange={(value) => handleChange('from_location_id', value)}>
          <SelectTrigger>
            <SelectValue placeholder="Select a location" />
          </SelectTrigger>
          <SelectContent>
            {locations.map(location => (
              <SelectItem key={location.id} value={location.id.toString()}>
                {location.name}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
      </div>

      <div className="space-y-2">
        <Label htmlFor="to_location_id">To Location</Label>
        <Select name="to_location_id" onValueChange={(value) => handleChange('to_location_id', value)}>
          <SelectTrigger>
            <SelectValue placeholder="Select a location" />
          </SelectTrigger>
          <SelectContent>
            {locations.map(location => (
              <SelectItem key={location.id} value={location.id.toString()}>
                {location.name}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
      </div>

      <Button type="submit" className="w-full">Request Issue</Button>

      {issueResult && (
        <div className={`mt-4 ${issueResult.includes('Error') ? 'text-red-500' : 'text-green-500'}`}>
          {issueResult}
        </div>
      )}
    </form>
  );
};

export default IssueForm;