'use client';

import React from 'react';
import Content from '@/components/Content';
import PortalCard from '@/components/PortalCard';
import { Space } from 'antd';
import { ExperimentFilled } from '@ant-design/icons';
import TestDetails from '@/components/TestDetails';

interface PageParams {
    params: {
        slug: string
    }
}

const Page: React.FC<PageParams> = ({ params }) => {
    const label = decodeURIComponent(atob(decodeURIComponent(params.slug)))
    return (
        <Content
            content={
                <Space direction="vertical" size="middle" style={{ display: 'flex' }}>
                    <PortalCard
                        icon={<ExperimentFilled />}
                        titleBits={[<span key="title">Test Details</span>]}
                    >
                        <TestDetails label={label} />
                    </PortalCard>
                </Space >
            }
        />
    );
}

export default Page;
