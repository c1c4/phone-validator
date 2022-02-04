import http from './api';
import IPhoneData from '../types/Phone';

const getAll = () => {
    return http.get<Array<IPhoneData>>('/phones');
};

const getAllWIthFilters = (filters: { [key: string]: unknown }) => {
    return http.get<Array<IPhoneData>>('/phones', { params: filters });
};

const PhoneService = {
    getAll,
    getAllWIthFilters,
};

export default PhoneService;
