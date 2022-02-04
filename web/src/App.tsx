import React, { ChangeEvent, useEffect, useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import IPhoneData from './types/Phone';
import PhoneService from './service/PhoneService';
import { AxiosResponse } from 'axios';

const App: React.FC = () => {
    const [phones, setPhones] = useState<IPhoneData[]>([]);
    const [country, setCountry] = useState<string>('');
    const [state, setState] = useState<string>('');
    const [page, setPage] = useState<number>(1);
    const [statePrev, setStatePrev] = useState<boolean>(false);
    const [stateNext, setStateNext] = useState<boolean>(false);

    const getPhones = () => {
        console.log('Passou aqui');
        const filter: { [key: string]: unknown } = {};

        if (country.length > 0 && state.length > 0) {
            filter['country'] = country;
            filter['state'] = state;
        } else if (country.length > 0) {
            filter['country'] = country;
        } else if (state.length > 0) {
            filter['state'] = state;
        }

        filter['page'] = page;

        PhoneService.getAllWIthFilters(filter)
            .then((response: AxiosResponse) => {
                setPhones(response.data.phones);
                setStatePrev(response.data.pagination.hasPrev);
                setStateNext(response.data.pagination.hasNext);
            })
            .catch((e: Error) => {
                console.log(e);
            });
    };

    const onChangeCountry = (e: ChangeEvent<HTMLSelectElement>) => {
        setCountry(e.target.value);
    };

    const onChangeState = (e: ChangeEvent<HTMLSelectElement>) => {
        setState(e.target.value);
    };

    const onClickPrev = () => {
        const newPage = page - 1;
        setPage(newPage);
    };

    const onClickNext = () => {
        const newPage = page + 1;
        setPage(newPage);
    };

    const divStyleApp = {
        paddingLeft: '20px',
    };

    const divStylePaddingTop = {
        paddingTop: '20px',
    };

    const divPhonesListTittle = {
        paddingTop: '20px',
        paddingBottom: '10px',
    };

    useEffect(getPhones, [country, page, state]);

    return (
        <div style={divStyleApp} id="app">
            <h1>Phone Numbers</h1>
            <div className="row" id="filters" style={divStylePaddingTop}>
                <div id="countryFilters" className="col-2">
                    <select className="form-select" value={country} onChange={onChangeCountry}>
                        <option value="">Select a country</option>
                        <option value="Cameroon">Cameroon</option>
                        <option value="Ethiopia">Ethiopia</option>
                        <option value="Morocco">Morocco</option>
                        <option value="Mozambique">Mozambique</option>
                        <option value="Uganda">Uganda</option>
                    </select>
                </div>
                <div id="stateFilters" className="col-3">
                    <select className="form-select" value={state} onChange={onChangeState}>
                        <option value="">Select a state</option>
                        <option value="OK">Valid phone numbers</option>
                        <option value="NOK">Invalid phone numbers</option>
                    </select>
                </div>
            </div>
            <div className="col" id="phones">
                <div className="row" id="phoneListTitle" style={divPhonesListTittle}>
                    <div className="col-3">Country</div>
                    <div className="col-3">State</div>
                    <div className="col-3">Country code</div>
                    <div className="col-3">Phone num</div>
                </div>
                {phones.map((phone) => (
                    <div className="row" key={phone.phoneNumberFormated}>
                        <div className="col-3">{phone.country}</div>
                        <div className="col-3">{phone.state}</div>
                        <div className="col-3">{phone.countryCode}</div>
                        <div className="col-3">{phone.phoneNumberFormated}</div>
                    </div>
                ))}
            </div>
            <div className="row justify-content-end" style={divStylePaddingTop}>
                <div className="col-1">
                    <button className="btn btn-secondary" disabled={!statePrev} onClick={onClickPrev}>
                        Previous
                    </button>
                </div>
                <div className="col-2">
                    <button className="btn btn-secondary" disabled={!stateNext} onClick={onClickNext}>
                        Next
                    </button>
                </div>
            </div>
        </div>
    );
};

export default App;
