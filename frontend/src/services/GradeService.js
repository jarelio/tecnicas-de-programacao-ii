import http from '../http-common';

const getAll = () => {
    return http.get('/grades');
};

const get = (id) => {
    return http.get(`/grades/${id}`);
};

const create = (data) => {
    return http.post('/grades', data);
};

const update = (id, data) => {
    return http.put(`/grades/${id}`, data);
};

const remove = (id) => {
    return http.delete(`/grades/${id}`);
};

const findByName = (student) => {
    return http.get(`/grades/student/${student}`);
};

export default {
    getAll,
    get,
    create,
    update,
    remove,
    findByName,
};
