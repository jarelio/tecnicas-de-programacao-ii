import React, { useState, useEffect } from 'react';
import GradeDataService from '../services/GradeService';
import { Link } from 'react-router-dom';

const GradeList = () => {
    const [grade, setGrade] = useState([]);
    const [currentGrade, setCurrentGrade] = useState(null);
    const [currentIndex, setCurrentIndex] = useState(-1);
    const [searchName, setSearchName] = useState('');

    useEffect(() => {
        retrieveGrade();
    }, []);

    const onChangeSearchName = (e) => {
        const searchName = e.target.value;
        setSearchName(searchName);
    };

    const retrieveGrade = () => {
        GradeDataService.getAll()
            .then((response) => {
                setGrade(response.data.result.data);
                console.log(response.data.result.data);
            })
            .catch((e) => {
                console.log(e);
            });
    };

    const setActiveGrade = (grade, index) => {
        setCurrentGrade(grade);
        setCurrentIndex(index);
    };

    const findByName = () => {
        GradeDataService.findByName(searchName)
            .then((response) => {
                setGrade(response.data.result.data);
                console.log(response.data.result.data);
            })
            .catch((e) => {
                console.log(e);
            });
    };

    return (
        <div className="list row">
            <div className="col-md-8">
                <div className="input-group mb-3">
                    <input
                        type="text"
                        className="form-control"
                        placeholder="Search by student"
                        value={searchName}
                        onChange={onChangeSearchName}
                    />
                    <div className="input-group-append">
                        <button
                            className="btn btn-outline-secondary"
                            type="button"
                            onClick={findByName}
                        >
                            Search
                        </button>
                    </div>
                </div>
            </div>
            <div className="col-md-6">
                <h4>Grade List</h4>

                <ul className="list-group">
                    {grade &&
                        grade.map(
                            (grade, index) =>
                                grade.student && (
                                    <li
                                        className={
                                            'list-group-item ' +
                                            (index === currentIndex
                                                ? 'active'
                                                : '')
                                        }
                                        onClick={() =>
                                            setActiveGrade(grade, index)
                                        }
                                        key={index}
                                    >
                                        {grade.student}
                                    </li>
                                )
                        )}
                </ul>
            </div>
            <div className="col-md-6">
                {currentGrade ? (
                    <div>
                        <h4>Grade</h4>
                        <div>
                            <label>
                                <strong>Student:</strong>
                            </label>{' '}
                            {currentGrade.student}
                        </div>
                        <div>
                            <label>
                                <strong>Subject:</strong>
                            </label>{' '}
                            {currentGrade.subject}
                        </div>
                        <div>
                            <label>
                                <strong>Type:</strong>
                            </label>{' '}
                            {currentGrade.type}
                        </div>
                        <div>
                            <label>
                                <strong>Value:</strong>
                            </label>{' '}
                            {currentGrade.value}
                        </div>

                        <Link
                            to={'/grade/' + currentGrade.id}
                            className="badge badge-warning"
                        >
                            Edit
                        </Link>
                    </div>
                ) : (
                    <div>
                        <br />
                        <p>Please click on a Grade...</p>
                    </div>
                )}
            </div>
        </div>
    );
};

export default GradeList;
