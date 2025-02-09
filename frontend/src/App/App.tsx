import React, { useEffect, useState } from 'react';
import { Table, Space } from 'antd';


interface Container {
    id: number;
    ipAddress: string;
    pingTime: number;
    lastSuccessful: string;
}

const App: React.FC = () => {
    const [containers, setContainers] = useState<Container[]>([]);

    useEffect(() => {
        const fetchContainers = async () => {
            const response = await fetch('http://backend:8080/containers');
            const data = await response.json();
            setContainers(data);
        };

        fetchContainers();
    }, []);

    return (
        <div>
            <h1>Container Status</h1>
            <Table dataSource={containers} pagination={false}>
                <Table.Column title="IP Address" dataIndex="ipAddress" key="ipAddress" />
                <Table.Column title="Ping Time (ms)" dataIndex="pingTime" key="pingTime" />
                <Table.Column title="Last Successful" dataIndex="lastSuccessful" key="lastSuccessful" />
            </Table>
        </div>
    );
};

export default App;